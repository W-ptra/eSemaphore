type chatLineProperty = {
    message: string;
    isSender: boolean;
}

function ChatLine({ message,isSender }: chatLineProperty) {

    const bgColor = isSender ? "bg-gray-100" : "bg-white"
    const chatLinePosition = isSender ? "justify-start" : "justify-end"

    return (
        <div
            className={`flex flex-row w-[100%] ${chatLinePosition}`}
        >
            <div
                className={`shadow flex p-2 m-2 justify-end rounded max-w-[90%] ${bgColor}`}
            >

                <div
                    className="mr-5"
                >
                    <p>
                        {message}
                    </p>
                </div>
                <div
                    className="text-xs items-end justify-end flex flex-col"
                >
                    <p>
                        11:11
                    </p>
                </div>
            </div>
        </div>
    )
}

export default ChatLine