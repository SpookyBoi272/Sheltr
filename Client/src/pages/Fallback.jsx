import { FiAlertCircle } from "react-icons/fi";

export default function Fallback() {
    return (
        <div className='flex-grow flex items-center justify-center flex-col'>
            <FiAlertCircle size={"120px"} />
            <h1 className='heading mt-6'>404 Page not Found</h1>
        </div>

    )
}