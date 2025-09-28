import axiosHelper from "./AxiosHelper"

export const getTrackRec = async (id : string) => {
    const res = await axiosHelper.get(`/songRec`)
    return res.data
}

export const getArtistRec = async (id : string) => {
    const res = await axiosHelper.get(`/artistRec`)
    return res.data
}

export const getAlbumRec = async (id : string) => {
    const res = await axiosHelper.get(`/albumRec`)
    return res.data
}