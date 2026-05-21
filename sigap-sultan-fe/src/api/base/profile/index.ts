import { baseFetcher } from "../fetcher";
import { apiProfile } from "../api-paths";
import { Profile } from "@/types/profile";

export interface UpdateProfileResponse {}

export interface UpdateProfileRequest extends Profile {}

class ProfileApi {
  getDetail() {
    return baseFetcher.apiGet<Profile>({
      path: `${apiProfile}`,
    });
  }
  update(request: UpdateProfileRequest) {
    return baseFetcher.apiPost<UpdateProfileResponse>({
      path: `${apiProfile}`,
      payload: request,
    });
  }
}

export const profileApi = new ProfileApi();
