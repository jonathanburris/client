// @flow

import * as Constants from '../constants/pgp'
import {safeTakeEvery} from '../util/saga'
import {pgpPgpStorageDismissRpc} from '../constants/types/flow-types'

import type {PgpAckedMessage} from '../constants/pgp'

function pgpStorageDismiss() {
  // make rpc call to pgpStorageDismiss
  pgpPgpStorageDismissRpc({
    callback: err => {
      if (err) {
        console.warn('Error in sending pgpPgpStorageDismissRpc:', err)
      }
    },
  })
}

const pgpSecurityModelChangeMessageSaga = function*({payload: {hitOk}}: PgpAckedMessage): any {
  pgpStorageDismiss()
}

const pgpSaga = function*(): any {
  yield safeTakeEvery(Constants.pgpAckedMessage, pgpSecurityModelChangeMessageSaga)
}

export default pgpSaga
