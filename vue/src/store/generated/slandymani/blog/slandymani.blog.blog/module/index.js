// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgCreateComment } from "./types/blog/tx";
import { MsgDeleteComment } from "./types/blog/tx";
import { MsgCreatePost } from "./types/blog/post";
import { MsgUpdateComment } from "./types/blog/tx";
const types = [
    ["/slandymani.blog.blog.MsgCreateComment", MsgCreateComment],
    ["/slandymani.blog.blog.MsgDeleteComment", MsgDeleteComment],
    ["/slandymani.blog.blog.MsgCreatePost", MsgCreatePost],
    ["/slandymani.blog.blog.MsgUpdateComment", MsgUpdateComment],
];
const registry = new Registry(types);
const defaultFee = {
    amount: [],
    gas: "200000",
};
const txClient = async (wallet, { addr: addr } = { addr: "http://localhost:26657" }) => {
    if (!wallet)
        throw new Error("wallet is required");
    const client = await SigningStargateClient.connectWithSigner(addr, wallet, { registry });
    const { address } = (await wallet.getAccounts())[0];
    return {
        signAndBroadcast: (msgs, { fee = defaultFee, memo = null }) => memo ? client.signAndBroadcast(address, msgs, fee, memo) : client.signAndBroadcast(address, msgs, fee),
        msgCreateComment: (data) => ({ typeUrl: "/slandymani.blog.blog.MsgCreateComment", value: data }),
        msgDeleteComment: (data) => ({ typeUrl: "/slandymani.blog.blog.MsgDeleteComment", value: data }),
        msgCreatePost: (data) => ({ typeUrl: "/slandymani.blog.blog.MsgCreatePost", value: data }),
        msgUpdateComment: (data) => ({ typeUrl: "/slandymani.blog.blog.MsgUpdateComment", value: data }),
    };
};
const queryClient = async ({ addr: addr } = { addr: "http://localhost:1317" }) => {
    return new Api({ baseUrl: addr });
};
export { txClient, queryClient, };
