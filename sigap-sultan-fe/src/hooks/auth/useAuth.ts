import { useContext } from "react";

import type { AuthContextType as JwtAuthContextType } from "@/contexts/auth";
import { AuthContext } from "@/contexts/auth";

type AuthContextType = JwtAuthContextType;

export const useAuth = <T = AuthContextType>() => useContext(AuthContext) as T;
