import { DefaultResponse, APIResponse } from "../../types/api";
import request, { AxiosPromise } from "axios";

export const NETWORK_ERROR = "NETWORK_ERROR";
export const GENERAL_ERROR = "GENERAL_ERROR";
export const BAD_REQUEST = "BAD_REQUEST";
export const UNAUTORIZED = "UNAUTORIZED";
export const NOT_FOUND = "NOT_FOUND";
export const FAILED_DEPENDENCY = "FAILED_DEPENDENCY";

export function createErrorMessage(httpCode: string, message: string): string {
  return `${httpCode ? httpCode + ": " : ""}${message}`;
}

export interface APIResult<T> {
  response?: APIResponse<T>;
  result?: T | null;
  error?: Error;
  errorCode?: string | null;
  errorData?: DefaultResponse;
  displayMessage?: string | null;
}

async function handleResponse<T>(promise: AxiosPromise): Promise<APIResult<T>> {
  try {
    const response = await promise;
    const responseData = response.data as APIResponse<T>;

    if (responseData.status !== 200) {
      return {
        error: new Error(createErrorMessage(String(response.status), "Error")),
        errorCode: String(responseData.status),
        errorData: responseData,
        displayMessage: responseData.message,
      };
    }

    return {
      response: responseData,
      result: responseData.data,
    };
  } catch (error: any) {
    if (request.isAxiosError(error) && error.response) {
      const response = error.response?.data as DefaultResponse;
      if (response) {
        return {
          error,
          errorCode: String(response.status),
          errorData: response,
          displayMessage: response.message,
        };
      }
    }

    console.error(`
      Error while trying to connect to 
      ${error}
    `);

    let errorCode = GENERAL_ERROR;

    if (error.message === "Network Error") {
      errorCode = NETWORK_ERROR;
    }

    return { error, errorCode };
  }
}

function generateErrorCode(httpCode: number) {
  switch (httpCode) {
    case 400: {
      return BAD_REQUEST;
    }
    case 401: {
      return UNAUTORIZED;
    }
    case 404: {
      return NOT_FOUND;
    }
    case 424: {
      return FAILED_DEPENDENCY;
    }
    default: {
      return GENERAL_ERROR;
    }
  }
}

export default handleResponse;
