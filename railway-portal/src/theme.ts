import { createTheme } from '@mui/material/styles';

const theme = createTheme({
  palette: {
    primary: {
      main: '#1abc9c', // Teal as the primary color
    },
    secondary: {
      main: '#f1c40f', // Navy blue as the secondary color
    },
  },
  typography: {
    fontFamily: '"Roboto", "Helvetica", "Arial", sans-serif',
    h4: {
      fontWeight: 600,
      color: '#009688',
    },
    h5: {
      fontWeight: 500,
    },
    button: {
      textTransform: 'none', // More natural button text
    },
  },
});

export default theme;
