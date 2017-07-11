import fetch from '../utils/fetch'

export interface Company {
    Name: string,
    City: string,
    Other: string
}

async function delay(ms:number) {
    const promise = new Promise((resolve, reject) => {
        setTimeout(() => {
            resolve(ms*10)
        }, 3000)
    })
    return await promise
}

export async function SetCompany(company: Company):Promise<Company> {
    let res = await fetch.post('company', company)
    let retc:Company = res.data
    return Promise.resolve(retc)
}

export async function GetCompany(nm:string):Promise<Company> {
    let res = await fetch.get('company', {params: {Name: nm}})
    let retc:Company = res.data
    return Promise.resolve(retc)
}
