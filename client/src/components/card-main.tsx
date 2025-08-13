import { Card, type CardProps } from '@mantine/core';
import React from 'react';

interface CardMainProps extends CardProps {
  children: React.ReactNode;
}

function CardMain({ children, ...props }: CardMainProps) {
  return (
    <Card withBorder radius="lg" bg="white" {...props}>
      {children}
    </Card>
  );
}

export default CardMain;
