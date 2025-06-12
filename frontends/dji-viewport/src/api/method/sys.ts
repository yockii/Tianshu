import type { CloudApiInfo, ConnectInfo } from "@/types/sys-info";
import { alovaInstance } from "../alova";

export const getCloudApiInfo = () => alovaInstance.Get<CloudApiInfo>('/sys/cloud-api-info')
export const getConnectInfo = () => alovaInstance.Get<ConnectInfo>('/sys/connect-info')