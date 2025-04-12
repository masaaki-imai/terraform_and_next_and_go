export type ValidationError<T> = {
  isValid: boolean;
  data: T | Record<string, string[]>; // Record<string, string[]>は、バリデーションエラーを表す
};
