
export const generateEndpointWithParam = (endpoint: string, params: Record<string, any>) => {
  let updatedEndpoint = endpoint;
  Object.entries(params).forEach(([key, value]) => {
    updatedEndpoint = updatedEndpoint.replace(`{${key}}`, value)
  });
  return updatedEndpoint;
}