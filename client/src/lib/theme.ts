import { generateColors } from '@mantine/colors-generator';
import { createTheme, type CSSVariablesResolver } from '@mantine/core';

export const cssVariablesResolver: CSSVariablesResolver = (theme) => ({
  variables: {},
  dark: {},
  light: {
    '--mantine-color-placeholder': theme.colors.neutral?.[4],
    '--mantine-color-dimmed': theme.colors.neutral?.[5],
    '--mantine-color-error': '#fe346e',
    '--mantine-color-body': '#f1f1f6ff',
  },
});

export const defaultTheme = createTheme({
  primaryColor: 'primary',
  defaultRadius: 'md',

  colors: {
    primary: generateColors('#263b44'),
  },
});
