import { styled, keyframes } from '../stitches.config';
import * as DialogPrimitive from '@radix-ui/react-dialog';

const overlayShow = keyframes({
  '0%': { opacity: 0 },
  '100%': { opacity: 1 },
});

const contentShow = keyframes({
  '0%': { opacity: 0, transform: 'translate(-50%, -48%) scale(.96)' },
  '100%': { opacity: 1, transform: 'translate(-50%, -50%) scale(1)' },
});

const Overlay = styled(DialogPrimitive.Overlay, {
  backgroundColor: '$blackA9',
  position: 'fixed',
  inset: 0,
  '@media (prefers-reduced-motion: no-preference)': {
    animation: `${overlayShow} 150ms cubic-bezier(0.16, 1, 0.3, 1)`,
  },
});

const Root = styled(DialogPrimitive.Root);
const Portal = styled(DialogPrimitive.Portal);
const Content = styled(DialogPrimitive.Content, {
  backgroundColor: '$slate1',
  position: 'fiixed',
  width: '400px',
  height: '300px',
  top: '50%',
  left: '50%',
});
const Trigger = styled(DialogPrimitive.Trigger);
const Title = styled(DialogPrimitive.Title);

export { Overlay, Content, Trigger, Portal, Title, Root };
