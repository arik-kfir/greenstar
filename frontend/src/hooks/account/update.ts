// This file is generated by greenstar scripts. DO NOT EDIT.
// noinspection DuplicatedCode

import { BaseAPIURL } from "../../services/util.ts"
import { Account } from "../../models/account.ts"
import { DateProperties } from "../../models/account.ts"
import { useOperation, Method, Hook } from "../../util/operation.ts"

export interface Request {
    id: string
    displayName: string
    icon?: string
    parentID?: string
}
export type Response = Account | undefined

export function useUpdateAccount(): Hook<Request, Response> {
    const opts = {
        initial: undefined,
        method: "PUT" as Method,
        url: (req: Request) => `${BaseAPIURL}/accounts/${req.id}`,
        dateProperties: DateProperties,
    }
    return useOperation<Request, Response>(opts)
}
