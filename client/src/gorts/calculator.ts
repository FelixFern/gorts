export type AdditionArgs = {
  A: number;
  B: number;
};

export type AdditionReply = number;

export type MultiplyArgs = {
  A: number;
  B: number;
};

export type MultiplyReply = number;

export interface Calculator {
  Addition(args: AdditionArgs): Promise<AdditionReply>;
  Multiply(args: MultiplyArgs): Promise<MultiplyReply>;
}