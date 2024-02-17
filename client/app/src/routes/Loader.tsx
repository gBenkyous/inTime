import React, { Suspense, ComponentType } from 'react';
import SuspenseLoader from 'src/components/SuspenseLoader';

const Loader = <P extends object>(Component: ComponentType<P>) => (props: P) =>
  (
    <Suspense fallback={<SuspenseLoader />}>
      <Component {...props} />
    </Suspense>
  );

export default Loader;
