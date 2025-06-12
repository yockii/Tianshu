
declare global {
  interface Window {
    thingConnectCallback?: (arg: any) => void;
    djiBridge: {
      platformVerifyLicense: (appId: string, appKey: string, appLicense: string) => string
      platformSetInformation: (platformName: string, title: string, desc: string) => string
      platformLoadComponent: (moduleName: string, param: string) => string
      apiSetToken: (token?: string) => string
      onBackClick: () => void
      onStopPlatform: () => void
    }
  }
}

interface JsResponse {
  code: number,
  message: string,
  data: any
}

export default {
  platformVerifyLicense(appId: string, appKey: string, appLicense: string): boolean {
    if (!window.djiBridge || !window.djiBridge.platformVerifyLicense) {
      console.error('djiBridge is not available')
      return false
    }
    return returnBool(window.djiBridge.platformVerifyLicense(appId, appKey, appLicense))
  },
  setPlatformMessage(platformName: string, workspaceName: string, desc: string): boolean {
    if (!window.djiBridge || !window.djiBridge.platformSetInformation) {
      console.error('djiBridge is not available')
      return false
    }
    const response = window.djiBridge.platformSetInformation(platformName, workspaceName, desc)
    return returnBool(response)
  },
  loadModule(moduleName: string, param: any): string {
    if (!window.djiBridge || !window.djiBridge.platformLoadComponent) {
      console.error(`djiBridge is not available`)
      return ''
    }
    const response = window.djiBridge.platformLoadComponent(moduleName, JSON.stringify(param))
    return returnString(response)
  },
  setToken(token?: string): string {
    if (!window.djiBridge || !window.djiBridge.apiSetToken) {
      console.error('djiBridge is not available')
      return ''
    }
    const response = window.djiBridge.apiSetToken(token)
    return returnString(response)
  },
  onBackClick(callback: () => void): void {
    if (!window.djiBridge || !window.djiBridge.onBackClick) {
      console.error('djiBridge is not available')
      return
    }
    window.djiBridge.onBackClick = callback
  },
  onStopPlatform(callback: () => void) {
    if (!window.djiBridge || !window.djiBridge.onBackClick) {
      console.error('djiBridge is not available')
      return
    }
    window.djiBridge.onStopPlatform = callback
  }
}

function returnString(response: string): string {
  const res: JsResponse = JSON.parse(response)
  return errorHint(res) ? res.data : ''
}

function returnBool(response: string): boolean {
  const res: JsResponse = JSON.parse(response)
  const isError = errorHint(res)
  if (JSON.stringify(res.data) !== '{}') {
    return isError && res.data
  }
  return isError
}

function errorHint(response: JsResponse): boolean {
  if (response.code !== 0) {
    console.error(response.message)
    return false
  }
  return true
}