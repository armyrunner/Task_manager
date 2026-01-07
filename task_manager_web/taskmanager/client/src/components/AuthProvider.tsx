import { useState } from 'react';
import type { ReactNode } from 'react';
import { AuthContext } from './AuthContext';

// Test user for development
const TEST_USER = {
  email: 'test@example.com',
  password: 'password123',
  name: 'Test User'
};

export function AuthProvider({ children }: { children: ReactNode }) {
  const [isLoggedIn, setIsLoggedIn] = useState(false);
  const [user, setUser] = useState<{ email: string; name: string } | null>(null);

  const login = (email: string, password: string): boolean => {
    // Check against test user
    if (email === TEST_USER.email && password === TEST_USER.password) {
      setIsLoggedIn(true);
      setUser({ email: TEST_USER.email, name: TEST_USER.name });
      return true;
    }
    return false;
  };

  const logout = () => {
    setIsLoggedIn(false);
    setUser(null);
  };

  return (
    <AuthContext.Provider value={{ isLoggedIn, user, login, logout }}>
      {children}
    </AuthContext.Provider>
  );
}
