import { createContext } from 'react';

export interface AuthContextType {
  isLoggedIn: boolean;
  user: { email: string; name: string } | null;
  login: (email: string, password: string) => boolean;
  logout: () => void;
}

export const AuthContext = createContext<AuthContextType | undefined>(undefined);
