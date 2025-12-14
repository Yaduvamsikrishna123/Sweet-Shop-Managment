import { render, screen, fireEvent } from '@testing-library/react';
import { describe, it, expect, vi } from 'vitest';
import { BrowserRouter } from 'react-router-dom';
import Login from '../pages/Login';
import { AuthContext } from '../context/AuthContext';

// Mock useNavigate
const mockNavigate = vi.fn();
vi.mock('react-router-dom', async () => {
    const actual = await vi.importActual('react-router-dom');
    return {
        ...actual,
        useNavigate: () => mockNavigate,
    };
});

describe('Login Page', () => {
    const mockLogin = vi.fn();

    const renderLogin = () => {
        render(
            <BrowserRouter>
                <AuthContext.Provider value={{ login: mockLogin }}>
                    <Login />
                </AuthContext.Provider>
            </BrowserRouter>
        );
    };

    it('renders login form', () => {
        renderLogin();
        expect(screen.getByRole('heading', { name: /welcome back/i })).toBeInTheDocument();
        expect(screen.getByPlaceholderText(/username/i)).toBeInTheDocument();
        expect(screen.getByPlaceholderText(/password/i)).toBeInTheDocument();
        expect(screen.getByRole('button', { name: /sign in/i })).toBeInTheDocument();
    });

    it('updates input fields', () => {
        renderLogin();
        const usernameInput = screen.getByPlaceholderText(/username/i);
        const passwordInput = screen.getByPlaceholderText(/password/i);

        fireEvent.change(usernameInput, { target: { value: 'testuser' } });
        fireEvent.change(passwordInput, { target: { value: 'password123' } });

        expect(usernameInput.value).toBe('testuser');
        expect(passwordInput.value).toBe('password123');
    });
});
