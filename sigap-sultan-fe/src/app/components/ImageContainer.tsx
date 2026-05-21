import React, { useState } from "react";
import Image, { ImageProps } from "next/image";
import { assetPrefix } from "@/utils/asset-prefix";

function ImageContainer(props: ImageProps) {
  const [error, setError] = useState(false);

  if (error) {
    return <Image {...props} src={assetPrefix("/error-image.png")} />;
  }

  return <Image {...props} onError={() => setError(true)} />;
}

export default ImageContainer;
