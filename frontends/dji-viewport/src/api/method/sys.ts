import type { CloudApiInfo } from "@/types/cloud-api";
import { alovaInstance } from "../alova";

export const getCloudApiInfo = () => alovaInstance.Get<CloudApiInfo>('/sys/cloud-api-info')