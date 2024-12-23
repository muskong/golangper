export default function access(initialState: { currentUser?: any }) {
  const { currentUser } = initialState || {};

  return {
    isAdmin: currentUser?.role === 'admin',
    isOperator: currentUser?.role === 'operator',
    canAdmin: currentUser?.role === 'admin' || currentUser?.role === 'operator',
  };
} 