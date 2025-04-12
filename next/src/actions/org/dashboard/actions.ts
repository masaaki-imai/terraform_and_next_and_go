'use server';

import { customFetch } from '@/src/utils/customFetch';

export interface LogoutResponse {
    isValid: boolean;
    result: boolean;
}

export async function logoutAction(): Promise<LogoutResponse> {
    return await customFetch<LogoutResponse>('/org/logout', {
        method: 'POST',
    });
}
