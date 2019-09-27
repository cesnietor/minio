/*
 * MinIO Cloud Storage (C) 2018 MinIO, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

import { createStore, applyMiddleware } from "redux"
import thunkMiddleware from "redux-thunk"
import reducers from "../reducers"
import { composeWithDevTools } from 'redux-devtools-extension';


const createStoreWithMiddleware = composeWithDevTools(applyMiddleware(thunkMiddleware))(createStore)

export default function configureStore(initialState) {
  const store = createStoreWithMiddleware(reducers, initialState)
  return store
}
