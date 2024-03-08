import service from '@/utils/request'

export const getWechatConfig = (data) => {
  return service({
    url: '/wechat/public/config',
    method: 'get',
    data
  })
}

export const updateWechatConfig = (data) => {
  return service({
    url: '/wechat/public/config',
    method: 'put',
    data
  })
}

