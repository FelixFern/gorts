import React, { useState } from "react";
import { Calculator } from "./calculator";

type RPCClientOptions = {
  serverUrl: string;
  setLoading: React.Dispatch<React.SetStateAction<boolean>>;
  setError: React.Dispatch<React.SetStateAction<string | null>>;
};

export const createRPCClient = <T extends Record<string, unknown>>({
  serverUrl,
  setLoading,
  setError,
}: RPCClientOptions): T => {
  const handleRPCCall = async (serviceName: string, methodName: string, args: unknown) => {
    setLoading(true);
    setError(null);

    try {
      const response = await fetch(serverUrl, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          id: Date.now(),
          jsonrpc: '2.0',
          method: `${serviceName}.${methodName}`,
          params: [args],
        }),
      });

      if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`);
      }

      const result = await response.json();

      if (result.error) {
        throw new Error(result.error.message || "RPC call failed");
      }

      return result.result;
    } catch (err) {
      const errorMessage = err instanceof Error ? err.message : "Unknown error occurred";
      setError(errorMessage);
      throw err;
    } finally {
      setLoading(false);
    }
  };

  return new Proxy({} as T, {
    get: (_, serviceName: string) =>
      new Proxy({}, {
        get: (_, methodName: string) =>
          (args: unknown) => handleRPCCall(serviceName, methodName, args),
      }),
  });
};


export const useRPCClient = () => {
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);

  const client = createRPCClient<{ Calculator: Calculator }>({
    serverUrl: "http://localhost:8080/rpc",
    setLoading,
    setError,
  });


  return {
    callRPC: client,
    loading,
    error,
  };
};
