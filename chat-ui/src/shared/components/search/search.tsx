import { useForm } from 'react-hook-form';
import styles from './search.module.scss';
import { ChangeEvent } from 'react';
import useDebounce from 'core/hooks/useDebounce';

interface FormInput {
  searchField: string;
}

interface SearchProps {
  sourceData: any[];
  excludedFields?: string[];
  handleSearchResult: (result: any[]) => void;
}

export default function Search({ sourceData, excludedFields, handleSearchResult }: SearchProps) {
  const { setData } = useDebounce({
    callback: (arg: any) => {
      if (!arg) {
        handleSearchResult(sourceData);
        return;
      }
      const matched = findData(arg, sourceData);
      handleSearchResult(matched);
    },
    debounce: 1000,
  });

  const {
    register,
    formState: { touchedFields, isValid, errors, isSubmitted },
  } = useForm<FormInput>({
    defaultValues: { searchField: '' },
    mode: 'onTouched', // default is onSubmit for validation to trigger
  });

  const handleOnChange = (event: ChangeEvent<HTMLInputElement>) => {
    event.preventDefault();
    setData(event.target.value);
  };

  const findData = (input: string, sourceData: any[]): any[] => {
    const regEx = new RegExp(input, 'i');
    const matched: any[] = [];

    sourceData.forEach(item => {
      if (!item) {
        // if null, to continue with next iteration
        return;
      }
      if (typeof item !== 'object') {
        // Primitive values.
        const result = regEx.test(item);
        if (result) {
          matched.push(item);
        }
      } else if (Array.isArray(item)) {
        const results = findData(input, item);
        matched.push(...results);
      } else {
        // item is an object.
        Object.keys(item).some(key => {
          if (!item[key] || (excludedFields && excludedFields.includes(key))) {
            return false;
          }

          if (typeof item[key] !== 'object') {
            // object value is primitive
            const result = regEx.test(item[key]);
            if (result) {
              // break out of key iteration and continue with next item
              matched.push(item);
              return true;
            }
          } else {
            // If value is an array or nested object, to call recursively with boolean result.
            // Only top-level objects are returned as results.
            const results = searchInNestedObject(input, item[key]);
            if (results) {
              matched.push(item);
              return true;
            }
          }
          return false;
        });
      }
    });
    return matched;
  };

  const searchInNestedObject = (input: string, arg: any) => {
    // Handler for iterating values of an object.
    // Arg can be an object or an array.
    // Will stop recursion if first value returns true.
    if (!arg) return false;

    const regEx = new RegExp(input, 'i');

    if (Array.isArray(arg)) {
      for (let i = 0; i < arg.length; i += 1) {
        if (typeof arg[0] !== 'object') {
          // iterating through primitive values.
          const result = regEx.test(arg[i]);
          if (result) {
            return true;
          }
        } else {
          // if nested value is an array, to call recursively.
          const result = searchInNestedObject(input, arg[i]);
          if (result) {
            return true;
          }
        }
      }
    } else {
      // nestedObj is an object
      Object.keys(arg).some(key => {
        if (typeof arg[key] !== 'object') {
          // iterating through primitive values
          const result = regEx.test(arg[key]);
          if (result) {
            return true;
          }
        } else {
          const result = searchInNestedObject(input, arg[key]);
          if (result) {
            return true;
          }
        }
        // No matched values in nested object values.
        return false;
      });
    }
    return false;
  };

  return (
    <div className={styles.wrapper}>
      <form className={styles.formWrapper}>
        <input
          {...register('searchField')}
          onChange={handleOnChange}
          id="search-input"
          autoComplete="on"
          className={`${styles.searchField} base-input`}
          placeholder="Search a user"></input>
      </form>
    </div>
  );
}
