import { AuthLoginResponse } from "@/api/base/auth";

interface UserAccountBusiness {
  businessId: string;
  createdAt: string;
  description: string;
  isActive: boolean;
  logo: string | null;
  name: string;
}

export interface UserAccount extends AuthLoginResponse {}
