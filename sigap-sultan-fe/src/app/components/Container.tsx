import React from 'react';

interface ContainerProps {
    children: React.ReactNode;
    className?: string;  // Adding an optional className prop
}

function Container({children, className}: ContainerProps) {
    return (
        <div className={`w-11/12 mx-auto mt-16 mb-16 ${className}`}>
            {children}
        </div>
    )
}

export {Container};
