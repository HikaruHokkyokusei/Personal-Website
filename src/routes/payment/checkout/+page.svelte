<script lang="ts">
    import { PUBLIC_SERVER_LOCATION_ORIGIN } from "$env/static/public";
    import { onMount } from "svelte";
    import { goto } from "$app/navigation";
    import { page } from "$app/stores";
    import axios from "axios";

    let showLoader = false, errorMessage = "", emailAddress = "", payableAmount;
    let stripe: App.IStripe, elements;

    onMount(() => {
        showLoader = true;
        const paymentId = $page.url.searchParams.get("pId");
        if (paymentId == null) {
            goto("/").then();
        } else {
            initialize(paymentId)
        }
    });

    const initialize = (paymentId) => {
        axios.get(`${PUBLIC_SERVER_LOCATION_ORIGIN}/payment/get-payment-intent/${encodeURI(paymentId)}`).then((response) => {
            if (response.status === 200) {
                let netAmount = response.data["netAmount"], decimals = netAmount % 100;
                payableAmount = `$${Math.floor(netAmount / 100)}.${decimals === 0 ? "00" : decimals}`;
                stripe = Stripe(response.data["publicKey"]);
                elements = stripe.elements({
                    "appearance": {
                        theme: 'night',
                    },
                    "clientSecret": response.data["clientSecret"]
                });

                const linkAuthenticationElement = elements.create("linkAuthentication");
                linkAuthenticationElement["mount"]("#link-authentication-element");
                linkAuthenticationElement.on('change', (event) => {
                    emailAddress = event.value["email"];
                });

                const paymentElement = elements.create("payment", {
                    layout: "tabs",
                });
                paymentElement["mount"]("#payment-element");

                showLoader = false;
            }
        }).catch((err) => {
            errorMessage = "This payment does not exists";
            console.log(err.response);
            setTimeout(() => {
                goto("/");
            }, 10000);
        });
    };

    const onPay = async () => {
        showLoader = true;
        errorMessage = "";

        const { error } = await stripe.confirmPayment({
            elements,
            redirect: "if_required",
            confirmParams: {},
        });

        if (error == null) {
            console.log("Payment Successful.");
        } else {
            if (error.type === "card_error" || error.type === "validation_error") {
                errorMessage = error.message;
            } else {
                errorMessage = "An unexpected error occurred.";
                console.log(error);
            }
            showLoader = false;
        }
    };
</script>

<div class="CenterRowFlex PaymentPageWrapper">
    <form class="CenterColumnFlex">
        <div id="link-authentication-element"></div>

        <div id="payment-element"></div>

        <div style="display: block;">
            <button class="CenterRowFlex PayNowButton" on:click|preventDefault={onPay} disabled={showLoader}>
                {#if showLoader}
                    <div class="LoadingHourglass"></div>
                {:else}
                    <span id="button-text">Pay {payableAmount}</span>
                {/if}
            </button>
            <div class="PaymentErrorMessageHolder">
                {errorMessage}
            </div>
        </div>
    </form>
</div>

<style>
    .PaymentPageWrapper {
        height: 100%;
        width: 100%;
    }

    form {
        width: 25%;
        min-width: 500px;
        aspect-ratio: 13 / 18;

        justify-content: space-around;
        align-items: normal;

        align-self: center;
        box-shadow: 0 0 0 1px rgba(50, 50, 93, 0.1), 0 2px 5px 0 rgba(50, 50, 93, 0.1), 0 1px 2px 0 rgba(0, 0, 0, 0.07);
        border-radius: 7px;
        padding: 40px 30px;
        background-color: rgba(255, 255, 255, 0.25);
    }

    .LoadingHourglass {
        height: 100%;
        display: inline-block;
    }

    .PayNowButton {
        height: 45px;
        width: 100%;

        box-sizing: border-box;
        padding: 12px 16px;

        background: #5469d4;
        color: #ffffff;

        cursor: pointer;
        font: 600 16px Arial, sans-serif;

        border-radius: 4px;
        box-shadow: 0 4px 5px 0 rgba(0, 0, 0, 0.1);

        transition: all 0.2s ease;
    }

    .PayNowButton:hover {
        filter: contrast(125%);
    }

    .PayNowButton:disabled {
        opacity: 0.5;
        cursor: default;
    }

    .LoadingHourglass:after {
        content: " ";
        height: 0;
        width: 0;

        display: block;
        box-sizing: border-box;
        margin: auto;

        border-radius: 50%;
        border: 10px solid;
        border-color: #fff transparent #fff transparent;

        animation: loading-hourglass 1.2s infinite;
    }

    .PaymentErrorMessageHolder {
        height: 25px;
        padding-top: 5px;

        color: #fe87a1;
        font-size: 16px;
        line-height: 20px;
        text-align: center;
    }

    @keyframes loading-hourglass {
        0% {
            transform: rotate(0);
            animation-timing-function: cubic-bezier(0.55, 0.055, 0.675, 0.19);
        }
        50% {
            transform: rotate(900deg);
            animation-timing-function: cubic-bezier(0.215, 0.61, 0.355, 1);
        }
        100% {
            transform: rotate(1800deg);
        }
    }
</style>