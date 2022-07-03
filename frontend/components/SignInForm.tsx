import { useFormik } from 'formik';

import useAuth from '../requests/useAuth';
import { signInFormSchema } from '../utils/validation';

const SignInForm = () => {
  const {
    signIn: { mutate, isLoading },
  } = useAuth();
  const { handleSubmit, handleChange, values, isValid, dirty } = useFormik({
    initialValues: {
      email: '',
      password: '',
    },
    onSubmit: (values) => {
      mutate(values);
    },
    validationSchema: signInFormSchema,
  });

  const isSubmitButtonDisabled = !isValid || !dirty || isLoading;

  return (
    <form onSubmit={handleSubmit}>
      <div>
        <div>
          <label htmlFor='email'>Work Email:</label>
          <input
            type='email'
            name='email'
            placeholder='Your work email'
            onChange={handleChange}
            value={values.email}
          />
        </div>
        <div>
          <label htmlFor='password'>Password</label>
          <input
            type='password'
            name='password'
            placeholder='*****'
            onChange={handleChange}
            value={values.password}
          />
        </div>
        <button type='submit' disabled={isSubmitButtonDisabled}>
          Sign in to toggler.in
        </button>
      </div>
    </form>
  );
};

export default SignInForm;
