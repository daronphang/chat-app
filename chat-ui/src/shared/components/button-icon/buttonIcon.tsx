import styles from './buttonIcon.module.scss';

interface Props {
  iconComponent: JSX.Element;
}

export default function ButtonIcon({ ...props }: Props) {
  return <button className={styles.buttonIcon}>{props.iconComponent}</button>;
}
