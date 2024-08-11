import { useAppDispatch } from 'core/redux/reduxHooks';
import { useForm } from 'react-hook-form';
import { useNavigate } from 'react-router-dom';
import styles from './newFriend.module.scss';

interface FormInput {
  email: string;
  displayName: string;
}

export default function NewFriend() {
  const dispatch = useAppDispatch();
  const navigate = useNavigate();
  const {
    register,
    getFieldState,
    handleSubmit,
    reset,
    formState: { touchedFields, isValid, errors, isSubmitted },
  } = useForm<FormInput>({
    defaultValues: { email: '', displayName: '' },
    mode: 'onTouched', // default is onSubmit for validation to trigger
  });

  const onSubmit = () => {};

  return (
    <div className={styles.wrapper}>
      <form className={styles.formWrapper}>
        <input {...register('email')} id="new-friend-email-input" autoComplete="on" className="input-field"></input>
      </form>
      <button className="btn-icon ms-3" onClick={handleSubmit(onSubmit)}>
        Submit
      </button>
    </div>
  );
}
