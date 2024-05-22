import request from '@/utils/myrequest'

export function getList(params) {
  return request({
    url: '/v1/deployments?ns=default',
    method: 'get',
    params
  })
}
