import 'remixicon/fonts/remixicon.css';
import { QueryClientProvider } from 'react-query';
import { ReactQueryDevtools } from 'react-query/devtools';

import type { AppProps } from 'next/app';
import { normalize } from 'normalize-stitches';
import { globalCss } from '../stitches.config';
import { queryClient } from '../utils/requestClient';

const globalStyles = globalCss({
  ...normalize,
  '@font-face': [
    {
      fontFamily: 'Open Sans',
      fontWeight: 400,
      src: 'url(fonts/OpenSans-Regular.ttf)',
    },
    {
      fontFamily: 'Open Sans',
      fontWeight: 500,
      src: 'url(fonts/OpenSans-Medium.ttf)',
    },
    {
      fontFamily: 'Open Sans',
      fontWeight: 600,
      src: 'url(fonts/OpenSans-SemiBold.ttf)',
    },
    {
      fontFamily: 'Open Sans',
      fontWeight: 700,
      src: 'url(fonts/OpenSans-Bold.ttf)',
    },
    {
      fontFamily: 'Open Sans',
      fontWeight: 400,
      fontStyle: 'italic',
      src: 'url(fonts/OpenSans-Italic.ttf)',
    },
  ],
  '*': {
    margin: 0,
    padding: 0,
    boxSizing: 'border-box',
    fontFamily: '$sans',
  },
  html: {
    width: '100%',
    height: '100%',
  },
  body: {
    width: '100%',
    height: '100%',
  },
});

function MyApp({ Component, pageProps }: AppProps) {
  globalStyles();
  return (
    <QueryClientProvider client={queryClient}>
      <Component {...pageProps} />
      <ReactQueryDevtools initialIsOpen={false} />
    </QueryClientProvider>
  );
}

export default MyApp;
