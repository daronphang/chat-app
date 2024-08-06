import { useAppSelector } from 'core/redux/reduxHooks';
import { Navigate, Outlet } from 'react-router-dom';

export const ProtectedRoutes = () => {
  const userId = useAppSelector(state => state.user.userId);
  return userId ? <Outlet /> : <Navigate to="/login" />;
};
