import { useFormik } from 'formik';
import type { NextPage } from 'next';
import useCreateUser from '../mutations/useCreateUser';

const Signup: NextPage = () => {
  const { handleCreateUser } = useCreateUser();
  const { handleChange, handleSubmit, values } = useFormik({
    initialValues: {
      email: '',
      password: '',
    },
    onSubmit: handleCreateUser,
  });

  return (
    <div>
      <form onSubmit={handleSubmit}>
        <label htmlFor='email'>
          Email:
          <input
            type='email'
            name='email'
            onChange={handleChange}
            value={values.email}
          />
        </label>
        <label htmlFor='password'>
          Password
          <input
            type='password'
            name='password'
            onChange={handleChange}
            value={values.password}
          />
        </label>
        <button type='submit'>Register</button>
      </form>
    </div>
  );
};

export default Signup;
