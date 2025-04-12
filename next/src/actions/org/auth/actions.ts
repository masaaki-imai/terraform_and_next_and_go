'use server'

import { customFetch } from '@/src/utils/customFetch'
import { ValidationResult } from '@/src/types/commons/validation'

type LoginParams = {
    email: string;
    password: string;
};

export const login = async (params: LoginParams) => {
    return await customFetch<ValidationResult<boolean>>('/org/login', {
        method: 'POST',
        body: JSON.stringify(params),
    });
};