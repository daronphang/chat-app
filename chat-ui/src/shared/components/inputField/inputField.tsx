import { forwardRef } from 'react';
import styles from './inputField.module.scss';
import { FieldError, FormState, UseFormGetFieldState } from 'react-hook-form';

// for input focus to target label, label must be AFTER input tag
// input:focus + label

// invalid and valid pseudo classes are used as react-hook-form
// cannot validate different input types i.e. email

interface InputProps {
  type?: string;
  id: string;
  label?: string;
  placeholder?: string;
  name?: string;
  isSubmitted?: boolean;
  maxLength?: number;
  styles?: any;
  fieldState: FieldState;
}

interface FieldState {
  invalid: boolean;
  isDirty: boolean;
  isTouched: boolean;
  isValidating: boolean;
  error?: FieldError | undefined;
}

const InputField = forwardRef(function renderInputField(props: InputProps, ref: React.ForwardedRef<any>) {
  //   const { isDirty, isTouched, invalid, error } = props.getFieldState;

  return (
    <div className={styles.wrapper}>
      <input
        className={styles.input}
        type={props.type}
        id={props.id}
        placeholder={props.placeholder}
        name={props.name}
        ref={ref}
        style={props.styles}
        maxLength={props.maxLength}
      />
      {props.label && <span className={styles.label}>{props.label}</span>}
    </div>
  );
});

export default InputField;
