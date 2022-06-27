import { styled } from '../stitches.config';
import Head from 'next/head';
import React from 'react';

const LayoutWrapper = styled('div', {
  paddingX: 'calc(6*$10)',
  paddingY: '$10',
  width: '100%',
  minHeight: '100vh',
  backgroundColor: '$slate1',

  '@bp1': {
    paddingX: 'calc(3*$10)',
  },

  '@bp2': {
    paddingX: '$10',
  },

  '@bp3': {
    paddingX: '$6',
  },
});

const Layout = ({ children }: { children: React.ReactNode }) => {
  return (
    <LayoutWrapper>
      <Head>
        <title>Toggler</title>
      </Head>
      {children}
    </LayoutWrapper>
  );
};

export default Layout;
