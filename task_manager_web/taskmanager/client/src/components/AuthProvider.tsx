import { useState } from 'react';
import type { ReactNode } from 'react';
import { AuthContext } from './AuthContext';



export function AuthProvider({ children }: { children: ReactNode }) {
  const [isLoggedIn, setIsLoggedIn] = useState(false);
  const [user, setUser] = useState<{id: number, username: string, email: string} | null>(null);

  useState(() => {
   const accessToken = localStorage.getItem('access_token');

   if (accessToken) {
    const userStoredInfo = localStorage.getItem('user');
    if (userStoredInfo) {
      try {
        const parsedUser = JSON.parse(userStoredInfo);
        setUser(parsedUser);
        setIsLoggedIn(true);
      } catch (error) {
        console.error('Error parsing user info:', error);
        localStorage.removeItem('user');
        localStorage.removeItem('access_token');
      }
    } else {
      console.warn('Access token found but no user data in local storage')
   }
  }
  });

  const login = (userData:{id: number, username: string, email: string}) => {
    setIsLoggedIn(true);
    setUser(userData);
    return true;
  };

  const logout = () => {
    localStorage.removeItem('user');
    localStorage.removeItem('access_token');
    localStorage.removeItem('refresh_token');
    setIsLoggedIn(false);
    setUser(null);
  };

  return (
    <AuthContext.Provider value={{ isLoggedIn, user, login, logout }}>
      {children}
    </AuthContext.Provider>
  );
}
