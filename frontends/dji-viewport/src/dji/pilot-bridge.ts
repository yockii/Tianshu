declare let window: {djiBridge: {
    platformVerifyLicense: (appId: string, appKey: string, appLicense: string) => string
}}
interface JsResponse{
  code:number,
  message:string,
  data:any
}

export default {
    platformVerifyLicense (appId:string, appKey:string, appLicense:string): boolean {
        if (!window.djiBridge || !window.djiBridge.platformVerifyLicense) {
            console.error('djiBridge is not available')
            return false
        }
        return returnBool(window.djiBridge.platformVerifyLicense(appId, appKey, appLicense))
    },
}

function returnBool (response: string): boolean {
  const res: JsResponse = JSON.parse(response)
  const isError = errorHint(res)
  if (JSON.stringify(res.data) !== '{}') {
    return isError && res.data
  }
  return isError
}

function errorHint (response: JsResponse): boolean {
  if (response.code !== 0) {
    console.error(response.message)
    return false
  }
  return true
}