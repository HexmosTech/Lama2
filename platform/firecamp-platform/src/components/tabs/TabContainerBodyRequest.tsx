import { FC, lazy, Suspense, useMemo, memo } from 'react';
import isEqual from 'react-fast-compare';
import _cloneDeep from 'lodash/cloneDeep';
import { ERequestTypes } from '@firecamp/types';
import { _object } from '@firecamp/utils';
import { Loader } from '@firecamp/ui';
import JsonTab from './requests/json/Request';
import MdTab from './requests/markdown/Request';

import { Rest } from '@firecamp/rest';
// import { GraphQL } from '@firecamp/graphql';
// import { WSClient } from '@firecamp/websocket';
// import { SocketIOClient } from '@firecamp/socket.io';

// const Rest = lazy(() =>
//   import('@firecamp/rest').then((module) => ({ default: module.Rest }))
// );


import { ETabEntityTypes, IEntityTab, IRequestTabProps } from './types';
import pltContext from '../../services/platform-context';
import { usePlatformStore } from '../../store/platform';
import EnvironmentTab from '../common/environment/tabs/Environment';
import CollectionFolderEntityTab from '../common/collection/CollectionFolderEntityTab';
import ImportTab from '../common/collection/ImportTab';

const TabContainerBodyRequest: FC<any> = ({ tab, index }) => {
  if (!tab || index === -1) {
    return <span />;
  }
  const { getFirecampAgent } = usePlatformStore.getState();

  if (
    [
      ERequestTypes.SocketIO,
      ERequestTypes.WebSocket,
      ERequestTypes.Rest,
      ERequestTypes.GraphQL,
    ].includes(tab.type)
  ) {
    tab = Object.assign({}, tab, {});
  }

  const tabProps: IRequestTabProps = useMemo(() => {
    return {
      tab,

      //v3 props
      platformContext: {
        ...pltContext,
        getFirecampAgent,
      },
    };
  }, [tab]);

  const _renderRequestTab = (tab: IEntityTab<any>) => {
    // console.log(tab.entity.__meta?.type, 'tab.entity.type');
    switch (tab.__meta.entityType) {
      case ETabEntityTypes.Request:
        const type = tab.entity?.__meta?.type;
        switch (type) {
          case ERequestTypes.Rest:
            return <Rest {...tabProps} />;
          // return (
          //   <Suspense fallback={<Loader />}>
          //     <Rest {...tabProps} />
          //   </Suspense>
          // );
           
          case 'json':
            return <JsonTab {...tabProps} />;
          case 'md':
            return <MdTab {...tabProps} />;
          default:
            return <span>Default Request Tab</span>;
        }
        break;
      case ETabEntityTypes.Environment:
        return <EnvironmentTab {...tabProps} />;
      case ETabEntityTypes.Collection:
      case ETabEntityTypes.Folder:
        return <CollectionFolderEntityTab {...tabProps} />;
      case ETabEntityTypes.Import:
        return <ImportTab {...tabProps} />;
      default:
        return <>No Entity Tab Found</>;
    }
  };

  return _renderRequestTab(tab);
};

export default memo(TabContainerBodyRequest, (pp, np) => {
  return isEqual(pp, np);
});
