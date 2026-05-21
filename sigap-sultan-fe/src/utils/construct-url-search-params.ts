
function constructUrlSearchParams<T>(payload: T): URLSearchParams {
  const formData = new URLSearchParams();
  if (payload) {
    for (const [key, value] of Object.entries(payload)) {
      // @ts-ignore
      if (typeof value !== 'undefined') formData.append(key, value);
    }
  }
  return formData;
}

export default constructUrlSearchParams;