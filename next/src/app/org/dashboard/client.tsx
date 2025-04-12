'use client';

import { usePathname, useRouter } from 'next/navigation';
import {
    Box,
    Container,
    Drawer,
    List,
    ListItem,
    ListItemButton,
    ListItemIcon,
    ListItemText,
    Typography,
    Divider,
    Dialog,
    DialogTitle,
    DialogContent,
    DialogContentText,
    DialogActions,
    Button,
} from '@mui/material';
import { People as PeopleIcon, Logout as LogoutIcon } from '@mui/icons-material';
import { logoutAction, LogoutResponse } from '@/src/actions/org/dashboard/actions';
import { useState } from 'react';

const DRAWER_WIDTH = 240;

const menuItems = [
    { text: 'アカウント一覧管理', icon: <PeopleIcon />, path: '/org/accounts' },
    { text: 'ユーザー招待', icon: <PeopleIcon />, path: '/org/users' },
];

export default function DashboardLayout() {
    const pathname = usePathname();
    const router = useRouter();
    const [isLogoutDialogOpen, setIsLogoutDialogOpen] = useState(false);

    const handleLogoutClick = () => {
        setIsLogoutDialogOpen(true);
    };

    const handleLogoutCancel = () => {
        setIsLogoutDialogOpen(false);
    };

    const handleLogoutConfirm = async () => {
        try {
            const logoutResponse: LogoutResponse = await logoutAction();
            if (logoutResponse.isValid && logoutResponse.result) {
                router.push('/public/org/login');
            } else {
                console.error('Logout response is not valid:', logoutResponse);
            }
        } catch (error) {
            console.error('Logout failed:', error);
        } finally {
            setIsLogoutDialogOpen(false);
        }
    };

    return (
        <Box sx={{ display: 'flex' }}>
            <Drawer
                variant="permanent"
                sx={{
                    width: DRAWER_WIDTH,
                    flexShrink: 0,
                    '& .MuiDrawer-paper': {
                        width: DRAWER_WIDTH,
                        boxSizing: 'border-box',
                        backgroundColor: '#ffffff',
                        borderRight: '1px solid rgba(0, 0, 0, 0.12)',
                    },
                }}
            >
                <Box sx={{ mt: 2 }}>
                    <List sx={{ width: '100%' }}>
                        {menuItems.map((item) => (
                            <ListItem key={item.text} disablePadding>
                                <ListItemButton
                                    selected={pathname === item.path}
                                    onClick={() => router.push(item.path)}
                                    sx={{
                                        pl: 3,
                                        '&.Mui-selected': {
                                            backgroundColor: '#e8f5e9',
                                            '&:hover': {
                                                backgroundColor: '#c8e6c9',
                                            },
                                        },
                                    }}
                                >
                                    <ListItemIcon sx={{ minWidth: 40, color: '#4caf50' }}>
                                        {item.icon}
                                    </ListItemIcon>
                                    <ListItemText
                                        primary={item.text}
                                        sx={{
                                            '& .MuiTypography-root': {
                                                fontSize: '0.9rem',
                                                fontWeight: pathname === item.path ? 'bold' : 'normal',
                                            },
                                        }}
                                    />
                                </ListItemButton>
                            </ListItem>
                        ))}
                        <Divider sx={{ my: 1 }} />
                        <ListItem disablePadding>
                            <ListItemButton
                                onClick={handleLogoutClick}
                                sx={{
                                    pl: 3,
                                    color: '#f44336',
                                    '&:hover': {
                                        backgroundColor: 'rgba(244, 67, 54, 0.08)',
                                    },
                                }}
                            >
                                <ListItemIcon sx={{ minWidth: 40, color: '#f44336' }}>
                                    <LogoutIcon />
                                </ListItemIcon>
                                <ListItemText
                                    primary="ログアウト"
                                    sx={{
                                        '& .MuiTypography-root': {
                                            fontSize: '0.9rem',
                                        },
                                    }}
                                />
                            </ListItemButton>
                        </ListItem>
                    </List>

                    <Dialog
                        open={isLogoutDialogOpen}
                        onClose={handleLogoutCancel}
                        aria-labelledby="logout-dialog-title"
                        aria-describedby="logout-dialog-description"
                    >
                        <DialogTitle id="logout-dialog-title">ログアウトの確認</DialogTitle>
                        <DialogContent>
                            <DialogContentText id="logout-dialog-description">
                                ログアウトしてよろしいですか？
                            </DialogContentText>
                        </DialogContent>
                        <DialogActions>
                            <Button onClick={handleLogoutCancel} color="primary">
                                キャンセル
                            </Button>
                            <Button onClick={handleLogoutConfirm} color="error" variant="contained">
                                ログアウト
                            </Button>
                        </DialogActions>
                    </Dialog>
                </Box>
            </Drawer>
            <Box
                component="main"
                sx={{
                    flexGrow: 1,
                    p: 3,
                    backgroundColor: '#f5f5f5',
                    minHeight: '100vh',
                }}
            >
                <Container maxWidth="lg">
                    <Box
                        sx={{
                            mb: 4,
                            display: 'flex',
                            justifyContent: 'space-between',
                            alignItems: 'center',
                            backgroundColor: '#ffffff',
                            padding: 2,
                            borderRadius: 1,
                            boxShadow: 1,
                        }}
                    >
                        <Box sx={{ display: 'flex', alignItems: 'center', gap: 2 }}>
                            <Typography variant="h5" component="h1" sx={{ fontWeight: 'bold' }}>
                                ダッシュボード
                            </Typography>
                        </Box>
                    </Box>
                </Container>
            </Box>
        </Box>
    );
}
