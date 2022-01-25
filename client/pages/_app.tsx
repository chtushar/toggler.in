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
      fontFamily: 'Noto Sans Display',
      fontWeight: 400,
      src: 'url(fonts/NotoSansDisplay-Regular.ttf)',
    },
    {
      fontFamily: 'Noto Sans Display',
      fontWeight: 500,
      src: 'url(fonts/NotoSansDisplay-Medium.ttf)',
    },
    {
      fontFamily: 'Noto Sans Display',
      fontWeight: 600,
      src: 'url(fonts/NotoSansDisplay-SemiBold.ttf)',
    },
    {
      fontFamily: 'Noto Sans Display',
      fontWeight: 700,
      src: 'url(fonts/NotoSansDisplay-Bold.ttf)',
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
