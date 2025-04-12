export type ValidationSuccess<T> = {
    isValid: true;
    result: T;
};

export type ValidationFailure = {
    isValid: false;
    result: ValidationError;
};

export type ValidationResult<T> = ValidationSuccess<T> | ValidationFailure;

export type ValidationError = string[];