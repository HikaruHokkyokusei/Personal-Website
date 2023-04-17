const commonProps = {
    "grouped": false,
    "barPercentage": 1,
    "borderRadius": 10,
    "borderSkipped": false,
    "borderWidth": 1
};

export const OdysseyMythos = {
    labels: ["Education", "Freelancing", "Job"],
    datasets: [
        {
            data: [
                { x: ["2019-06-25", "2023-05-01"], y: "Education", title: "KIIT D. University", content: "" }
            ],
            backgroundColor: "rgb(253, 98, 131, 0.2)",
            borderColor: "rgb(253, 98, 131, 1)",
            ...commonProps
        },
        {
            data: [
                { x: ["2023-08-27", "2025-06-01"], y: "Education", title: "Stony Brook University", content: "" }
            ],
            backgroundColor: "rgb(253, 98, 131, 0.2)",
            borderColor: "rgb(253, 98, 131, 1)",
            ...commonProps
        },
        {
            data: [
                { x: ["2020-08-01", "2022-06-01"], y: "Freelancing", title: "Upwork Global", content: "" }
            ],
            backgroundColor: "rgb(74, 190, 190, 0.2)",
            borderColor: "rgb(74, 190, 190, 1)",
            ...commonProps
        },
        {
            data: [
                { x: ["2021-10-01", "2022-03-01"], y: "Job", title: "Highradius", content: "" }
            ],
            backgroundColor: "rgb(253, 158, 63, 0.2)",
            borderColor: "rgb(253, 158, 63, 1)",
            ...commonProps
        },
        {
            data: [
                { x: ["2022-06-06", "2023-07-01"], y: "Job", title: "Autodesk India Pvt. Ltd.", content: "" }
            ],
            backgroundColor: "rgb(253, 158, 63, 0.2)",
            borderColor: "rgb(253, 158, 63, 1)",
            ...commonProps
        }
    ]
};