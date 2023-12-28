import _cleanDeep from 'clean-deep';
import _cloneDeep from 'lodash/cloneDeep';
import {
  EAuthTypes,
  EScriptLanguages,
  EScriptTypes,
  IAuth,
  IRest,
  IScript,
} from '@firecamp/types';
import { _clipboard } from '@firecamp/utils';
import { _object } from '@firecamp/utils';
import { TStoreSlice } from '../store.type';

interface IExecutionSlice {
  /**
   * prepare auth for execution
   * return the IAuth from request or from parent artifacts
   * 1. if request has auth set then return request auth
   * 2. if type is inherit then return parent auth
   *  2.1 if folder has auth then return folder auth
   *  2.2 else return collection auth
   */
  prepareAuthForExecution: () => IAuth;
  /**
   * prepare preScripts & postScripts for execution
   *
   * return the merged/combined pre/post scripts of request and parent (folder/collection)
   * then merging sequence will be for
   * 1. preScripts
   *    sequence collection + folder + request
   * 2. postScripts
   *    sequence collection + folder + request
   */
  prepareScriptsForExecution: () => {
    preScripts: IScript[];
    postScripts: IScript[];
  };
  /**
   * return final request for execution,
   *
   * the request will be generated based on the configuration with parent artifacts,
   * here the auth, preScripts & postScripts might get manipulated by the parent artifacts at the time of execution
   */
  prepareRequestForExecution: () => IRest;
  execute(): void;
}

const findScriptValue = (
  scripts: IScript[],
  type: EScriptTypes
): IScript['value'] => {
  const s = scripts.find((s) => s.type == type);
  return s?.value || [];
};

const createExecutionSlice: TStoreSlice<IExecutionSlice> = (set, get) => ({
  prepareAuthForExecution: () => {
    const {
      runtime: { parentArtifacts },
      request: { auth },
    } = get();
    if (auth?.type != EAuthTypes.Inherit) return auth;
    if (!parentArtifacts) return { type: EAuthTypes.None, value: '' };
    const { collection, folder } = parentArtifacts;
    if (folder?.auth?.type) return folder.auth;
    else if (collection?.auth?.type) return collection.auth;
    else return { type: EAuthTypes.None, value: '' };
  },
  prepareScriptsForExecution: () => {
    const {
      runtime: { parentArtifacts },
      request,
    } = get();
    const { collection, folder } = parentArtifacts;
    const scripts = {
      request: { pre: [], post: [] },
      collection: { pre: [], post: [] },
      folder: { pre: [], post: [] },
    };

    if (request?.preScripts?.length) {
      scripts.request.pre = findScriptValue(
        request.preScripts,
        EScriptTypes.PreRequest
      );
    }
    if (request?.postScripts?.length) {
      scripts.request.post = findScriptValue(
        request.postScripts,
        EScriptTypes.Test
      );
    }

    if (collection?.preScripts?.length) {
      scripts.collection.pre = findScriptValue(
        collection.preScripts,
        EScriptTypes.PreRequest
      );
    }
    if (collection?.postScripts?.length) {
      scripts.collection.post = findScriptValue(
        collection.postScripts,
        EScriptTypes.Test
      );
    }

    if (folder?.preScripts?.length) {
      scripts.folder.pre = findScriptValue(
        folder.preScripts,
        EScriptTypes.PreRequest
      );
    }
    if (folder?.postScripts?.length) {
      scripts.folder.post = findScriptValue(
        folder.postScripts,
        EScriptTypes.Test
      );
    }

    return {
      preScripts: [
        {
          id: 'pre-script',
          value: [
            ...scripts.collection.pre,
            ...scripts.folder.pre,
            ...scripts.request.pre,
          ],
          type: EScriptTypes.PreRequest,
          language: EScriptLanguages.JavaScript,
        },
      ],
      postScripts: [
        {
          id: 'tests',
          value: [
            ...scripts.collection.post,
            ...scripts.folder.post,
            ...scripts.request.post,
          ],
          type: EScriptTypes.Test,
          language: EScriptLanguages.JavaScript,
        },
      ],
    };
  },

  prepareRequestForExecution: () => {
    const { request, prepareAuthForExecution, prepareScriptsForExecution } =
      get();
    const auth = prepareAuthForExecution();
    const { preScripts, postScripts } = prepareScriptsForExecution();
    return {
      ...request,
      auth,
      preScripts,
      postScripts,
    };
  },
  execute: async () => {
    const state = get();
    // set response empty
    set({ response: { code: 0 } });

    // Check if request is running or not. stop running request if already true
    if (state.runtime.isRequestRunning === true) {
      await state.context.request.cancelExecution(state.request.__ref.id);
      // set request running state as false
      state.setRequestRunningFlag(false);
      return;
    }
    state.setRequestRunningFlag(true);

    const finalRequest = state.prepareRequestForExecution();
    console.log(finalRequest, '...finalRequest');
    let response = {
      code: 200,
      status: '',
      body: '{\n  "args": {}, \n  "headers": {\n    "Accept": "*/*", \n    "Accept-Encoding": "gzip, deflate, br", \n    "Accept-Language": "en-GB,en-US;q=0.9,en;q=0.8", \n    "Content-Type": "application/json", \n    "Host": "httpbin.org", \n    "Sec-Ch-Ua": "\\"Google Chrome\\";v=\\"119\\", \\"Chromium\\";v=\\"119\\", \\"Not?A_Brand\\";v=\\"24\\"", \n    "Sec-Ch-Ua-Mobile": "?0", \n    "Sec-Ch-Ua-Platform": "\\"Linux\\"", \n    "Sec-Fetch-Dest": "empty", \n    "Sec-Fetch-Mode": "cors", \n    "Sec-Fetch-Site": "cross-site", \n    "User-Agent": "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/119.0.0.0 Safari/537.36", \n    "X-Amzn-Trace-Id": "Root=1-6568d62e-32fe0ea5560563e15fbc98bb"\n  }, \n  "origin": "103.169.215.32", \n  "url": "https://httpbin.org/get"\n}\n',
      headers: [
        {
          key: 'access-control-allow-credentials',
          value: 'true',
          type: 'text',
          disable: false,
          description: '',
        },
        {
          key: 'access-control-allow-origin',
          value: '*',
          type: 'text',
          disable: false,
          description: '',
        },
        {
          key: 'content-length',
          value: '745',
          type: 'text',
          disable: false,
          description: '',
        },
        {
          key: 'content-type',
          value: 'application/json',
          type: 'text',
          disable: false,
          description: '',
        },
        {
          key: 'date',
          value: 'Thu, 30 Nov 2023 18:36:29 GMT',
          type: 'text',
          disable: false,
          description: '',
        },
        {
          key: 'server',
          value: 'gunicorn/19.9.0',
          type: 'text',
          disable: false,
          description: '',
        },
      ],
      responseTime: 975,
      responseSize: 745,
      timeline:
        '\n-----------   GENERAL  -----------\n\n# Request URL:  https://httpbin.org/get\n# Request Method: get\n# Status Code: 200 \n\n-----------   RESPONSE HEADERS   -----------\n\n< access-control-allow-credentials:true\n< access-control-allow-origin:*\n< content-length:745\n< content-type:application/json\n< date:Thu, 30 Nov 2023 18:36:29 GMT\n< server:gunicorn/19.9.0\n\n\n-----------   RESPONSE DATA   -----------\n\n> {\n  "args": {}, \n  "headers": {\n    "Accept": "application/json, text/plain, */*", \n    "Accept-Encoding": "gzip, deflate, br", \n    "Accept-Language": "en-GB,en-US;q=0.9,en;q=0.8", \n    "Host": "httpbin.org", \n    "Sec-Ch-Ua": "\\"Google Chrome\\";v=\\"119\\", \\"Chromium\\";v=\\"119\\", \\"Not?A_Brand\\";v=\\"24\\"", \n    "Sec-Ch-Ua-Mobile": "?0", \n    "Sec-Ch-Ua-Platform": "\\"Linux\\"", \n    "Sec-Fetch-Dest": "empty", \n    "Sec-Fetch-Mode": "cors", \n    "Sec-Fetch-Site": "cross-site", \n    "User-Agent": "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/119.0.0.0 Safari/537.36", \n    "X-Amzn-Trace-Id": "Root=1-6568d62d-6cf2d0257610f2b25eed0f96"\n  }, \n  "origin": "103.169.215.32", \n  "url": "https://httpbin.org/get"\n}\n',
      cookies: [],
    };

    let convertedString = '';

    if (
      typeof finalRequest.body.value === 'string' &&
      finalRequest.body.value !== ''
    ) {
      let jsonObject = JSON.parse(finalRequest.body.value);

      convertedString = Object.entries(jsonObject)

        .map(([key, value]) => `${key}='${value}'`)

        .join('\n');
    }

    console.log('first-string', convertedString);

    const headersList = finalRequest.headers;

    let headers = headersList

      .filter((item) => !item.disable) // Filter out disabled items

      .map((item) => `${item.key}:'${item.value}'`) // Map to 'key:value' format

      .join('\n'); // Join all items with newline character
    console.log('headers', headers);
    let command = '';
    // if (fcRequest.method == 'POST')
    command = `${finalRequest.method}\n${finalRequest.url.raw}\n${finalRequest.body.value}\n  ${headers}\n`;
    console.log('command', command);
    // else command = `${fcRequest.method}\n${fcRequest.url.raw}`;
    const lama2req = await window?.makeLamaRequest?.(command);

    console.log('lama2req', lama2req);

    response['body'] = lama2req;
    if (response) {
      set((s) => ({ response }));
      // TODO: check what to set/ response or testScriptResponse
    }

    state.setRequestRunningFlag(false);

    // await state.context.request
    //   .execute(finalRequest)
    //   .then(({ response, variables, testResult, scriptErrors }) => {
    //     // console.log({ response, variables, testResult });
    //     if (response?.error) {
    //       const error = response.error;
    //       // console.log(error.message, error.code, error.e.response, error.e);
    //     }
    //     if (response) {
    //       console.log('response', { response });
    //       set((s) => ({ response, testResult, scriptErrors })); // TODO: check what to set/ response or testScriptResponse
    //     }
    //   })
    //   .catch((e) => {
    //     console.log('debug-1', e.message, e.stack, e.response, e, 9090);
    //   })
    //   .finally(() => {
    //     state.setRequestRunningFlag(false);
    //   });
    // execute request
    //     let convertedString = '';

    //     if (
    //       typeof finalRequest.body.value === 'string' &&
    //       finalRequest.body.value !== ''
    //     ) {
    //       let jsonObject = JSON.parse(finalRequest.body.value);

    //       convertedString = Object.entries(jsonObject)

    //         .map(([key, value]) => `${key}='${value}'`)

    //         .join('\n');
    //     }

    //     console.log('first-string', convertedString);
    //     let response = {};
    //     const headersList = finalRequest.headers;

    //     let headers = headersList

    //       .filter((item) => !item.disable) // Filter out disabled items

    //       .map((item) => `${item.key}:'${item.value}'`) // Map to 'key:value' format

    //       .join('\n'); // Join all items with newline character
    //     console.log('headers', headers);
    //     let command = '';
    //     // if (fcRequest.method == 'POST')
    //     command = `${finalRequest.method}\n${finalRequest.url.raw}\n${finalRequest.body.value}\n  ${headers}\n`;
    //     console.log('command', command);
    //     // else command = `${fcRequest.method}\n${fcRequest.url.raw}`;
    //     const lama2req = await window?.makeLamaRequest?.(command);

    //     console.log('lama2req', lama2req);

    //     response['body'] = lama2req;
    //     if (response) {
    //       set((s) => ({ response }));
    //       // TODO: check what to set/ response or testScriptResponse
    //     }
    //     state.setRequestRunningFlag(false);
    //     console.log('response', response);
  },
});

export { createExecutionSlice, IExecutionSlice };
