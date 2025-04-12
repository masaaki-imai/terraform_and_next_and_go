'use client';

import { useState } from 'react';
import { useRouter } from 'next/navigation';
import { login } from '@/src/actions/org/auth/actions';
import ValidationModal from '@/src/components/organisms/org/validationModal';
import {
    Box,
    Container,
    Typography,
    TextField,
    Button,
    Paper,
    Link,
} from '@mui/material';

export default function LoginFormContainer() {
    const router = useRouter();
    const [formValues, setFormValues] = useState({ email: '', password: '' });
    const [errorMessages, setErrorMessages] = useState<string[]>([]);
    const [isOpen, setIsOpen] = useState(false);

    const handleInputChange = (name: string, value: string) => {
        setFormValues((prev) => ({ ...prev, [name]: value }));
    };

    const handleSubmit = async () => {
        const result = await login({
            email: formValues.email,
            password: formValues.password,
        });
        if (!result.isValid) {
            setErrorMessages(result.result);
            setIsOpen(true);
        } else if (result.result) {
            router.push('/org/dashboard');
        }
    };

    return (
        <Box
            sx={{
                minHeight: '100vh',
                backgroundColor: '#ffffff',
                display: 'flex',
                flexDirection: 'column',
                alignItems: 'center',
                pt: 8,
                pb: 8,
                fontFamily:
                    '-apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, "Helvetica Neue", Arial, sans-serif',
            }}
        >
            <Container component="main" maxWidth="sm">
                <Box sx={{ display: 'flex', flexDirection: 'column', alignItems: 'center' }}>
                    <Typography
                        component="h1"
                        variant="h4"
                        sx={{ mb: 4, fontWeight: 'bold', color: '#333333' }}
                    >
                        ログイン
                    </Typography>
                    <Paper sx={{ width: '100%', p: 4, borderRadius: 2 }}>
                        <TextField
                            margin="normal"
                            fullWidth
                            id="email"
                            label="メールアドレス"
                            name="email"
                            autoComplete="email"
                            autoFocus
                            value={formValues.email}
                            onChange={(e) => handleInputChange(e.target.name, e.target.value)}
                            sx={{
                                mb: 3,
                                '& .MuiOutlinedInput-root': {
                                    backgroundColor: '#f8f9fa',
                                    '&:hover fieldset': { borderColor: '#01c200' },
                                    '&.Mui-focused fieldset': { borderColor: '#01c200' },
                                },
                            }}
                        />
                        <TextField
                            margin="normal"
                            fullWidth
                            name="password"
                            label="パスワード"
                            type="password"
                            id="password"
                            autoComplete="current-password"
                            value={formValues.password}
                            onChange={(e) => handleInputChange(e.target.name, e.target.value)}
                            sx={{
                                mb: 2,
                                '& .MuiOutlinedInput-root': {
                                    backgroundColor: '#f8f9fa',
                                    '&:hover fieldset': { borderColor: '#01c200' },
                                    '&.Mui-focused fieldset': { borderColor: '#01c200' },
                                },
                            }}
                        />

                        <Button
                            type="button"
                            fullWidth
                            variant="contained"
                            onClick={handleSubmit}
                            sx={{
                                mt: 3,
                                mb: 2,
                                py: 1.8,
                                fontSize: '1rem',
                                fontWeight: 500,
                                backgroundColor: '#01c200',
                                color: '#ffffff',
                                textTransform: 'none',
                                borderRadius: '4px',
                                boxShadow: 'none',
                                letterSpacing: '0.5px',
                                transition: 'all 0.3s ease',
                                '&:hover': {
                                    backgroundColor: '#00a000',
                                    transform: 'translateY(-1px)',
                                    boxShadow: '0 2px 4px rgba(1, 194, 0, 0.2)',
                                },
                                '&:active': {
                                    transform: 'translateY(0)',
                                    boxShadow: 'none',
                                    backgroundColor: '#019800',
                                },
                                '&.Mui-disabled': {
                                    backgroundColor: '#7dd87c',
                                    color: '#ffffff',
                                },
                            }}
                        >
                            ログイン
                        </Button>

                        <Box sx={{ mt: 2, display: 'flex', justifyContent: 'center', gap: 10 }}>
                            <Link
                                href="/public/org/signup"
                                variant="body2"
                                sx={{
                                    color: '#01c200',
                                    textDecoration: 'none',
                                    fontWeight: 500,
                                    transition: 'all 0.3s ease',
                                    '&:hover': {
                                        color: '#00a000',
                                        textDecoration: 'none',
                                        opacity: 0.8,
                                    },
                                }}
                            >
                                新規登録
                            </Link>
                            <Link
                                href="/public/org/forgetPassword"
                                variant="body2"
                                sx={{
                                    color: '#01c200',
                                    textDecoration: 'none',
                                    fontWeight: 500,
                                    transition: 'all 0.3s ease',
                                    '&:hover': {
                                        color: '#00a000',
                                        textDecoration: 'none',
                                        opacity: 0.8,
                                    },
                                }}
                            >
                                パスワードを忘れた方はこちら
                            </Link>
                        </Box>
                    </Paper>
                </Box>
            </Container>
            <ValidationModal open={isOpen} onClose={() => setIsOpen(false)} errors={errorMessages} />
        </Box>
    );
}

