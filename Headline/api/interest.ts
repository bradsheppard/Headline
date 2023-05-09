import {InterestServiceClient} from "../proto/interest/InterestServiceClientPb";
import {GetInterestsRequest} from "../proto/interest/interest_pb";
import {API_HOST} from "./constants";

class InterestService {
    static async getInterests(userId: number): Promise<Array<string>> {
        const interestServiceClient = new InterestServiceClient(`http://${API_HOST}:80`)

        const request = new GetInterestsRequest()
        request.setUserid(userId)

        const response = await interestServiceClient.getInterests(request, null)
        return response.getInterestsList().map(x => x.getName())
    }
}

export default InterestService;

