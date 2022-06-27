import * as yup from 'yup';

export const signupFormSchema = yup.object().shape({
  name: yup.string().required(),
  email: yup.string().email().required(),
  password: yup.string().min(8).required(),
});

export const signInFormSchema = yup.object().shape({
  email: yup.string().email().required(),
  password: yup.string().min(8).required(),
});
