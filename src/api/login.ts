import fetch from '../utils/fetch'

export interface UserAuth{
    Id: string,
    Password: string
}

export async function LoginWithAccount(u:UserAuth) {
    let ret
    try {
        let res = await fetch.post('login', u)
        console.log('login response=', res)
        ret = res.data
    } catch (err) {
        console.log('login err=', err)
        ret=null
    }
    return Promise.resolve(ret)
}

export async function Logout() {
    let res = await fetch.post('logout')
    console.log('logout response=', res)
    let ret = res.data
    return Promise.resolve(ret)
}

