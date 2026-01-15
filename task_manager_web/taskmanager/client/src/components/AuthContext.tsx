import { createContext } from 'react';

export interface AuthContextType {
  isLoggedIn: boolean;
  user: {id: number, username: string, email: string} | null;
  login: (userData:{id: number, username: string, email: string}) => void;
  logout: () => void;
}

export const AuthContext = createContext<AuthContextType | undefined>(undefined);
