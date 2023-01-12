import axios from '@/libs/api.request'

export const createUser = ({ userName, password }) => {
  const data = {
    userName,
    password
  }
  return axios.request({
    url: '/api/v1/users',
    data,
    method: 'post'
  })
}
