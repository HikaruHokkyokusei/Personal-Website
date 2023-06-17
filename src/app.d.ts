import type { default as Web3 } from "web3";

// See https://kit.svelte.dev/docs/types#app
// for information about these interfaces
declare global {
    namespace App {
        // interface Error {}
        // interface Locals {}
        // interface PageData {}
        // interface Platform {}
        interface IStripe {
            [key: string]: any
        }
    }

    // noinspection JSUnusedGlobalSymbols
    interface Window {
        ethereum: any;
        web3: Web3;
    }

    const Stripe: (string) => App.IStripe;
}

export {};
