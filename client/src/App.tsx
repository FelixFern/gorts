import { useState } from 'react'
import './App.css'
import reactLogo from './assets/react.svg'
import { useRPCClient } from './gorts/useRPCClient'

function App() {
  const [variableA, setVariableA] = useState(0)
  const [variableB, setVariableB] = useState(0)
  const [result, setResult] = useState<number | null>(null)
  const { callRPC } = useRPCClient()

  const handleCalculate = async () => {
    const res = await callRPC.Calculator.Multiply({
      A: variableA,
      B: variableB
    })
    setResult(res)
  }

  return (
    <>
      <div>
        <a href="https://react.dev" target="_blank">
          <img src={reactLogo} className="logo react" alt="React logo" />
        </a>
      </div>
      <h1>gorts Client</h1>
      <div className='card'>
        <label>Variable A: </label>
        <input placeholder='Variable A' onChange={(e) => setVariableA(Number(e.currentTarget.value))}></input>
      </div>
      <div className='card'>
        <label >Variable B: </label>
        <input placeholder='Variable B' onChange={(e) => setVariableB(Number(e.currentTarget.value))}></input>
      </div>
      <button onClick={() => handleCalculate()}>Calculate</button>
      <div>
        Result: {result}
      </div>
    </>
  )
}

export default App
