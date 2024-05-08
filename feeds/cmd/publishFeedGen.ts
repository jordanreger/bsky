import "https://deno.land/std/dotenv/load.ts";

import { AtpAgent, BlobRef } from "npm:@atproto/api";

const states = [
  "al",
  "ak",
  "az",
  "ar",
  "as",
  "ca",
  "co",
  "ct",
  "de",
  "dc",
  "fl",
  "ga",
  "gu",
  "hi",
  "id",
  "il",
  "in",
  "ia",
  "ks",
  "ky",
  "la",
  "me",
  "md",
  "ma",
  "mi",
  "mn",
  "ms",
  "mo",
  "mt",
  "ne",
  "nv",
  "nh",
  "nj",
  "nm",
  "ny",
  "nc",
  "nd",
  "mp",
  "oh",
  "ok",
  "or",
  "pa",
  "pr",
  "ri",
  "sc",
  "sd",
  "tn",
  "tx",
  "tt",
  "ut",
  "vt",
  "va",
  "vi",
  "wa",
  "wv",
  "wi",
  "wy",
];

const run = async (state: string) => {
  // YOUR bluesky handle
  // Ex: user.bsky.social
  const handle = "jordanreger.com";

  // YOUR bluesky password, or preferably an App Password (found in your client settings)
  // Ex: abcd-1234-efgh-5678
  const password = Deno.env.get("bskypass")!;

  // A short name for the record that will show in urls
  // Lowercase with no spaces.
  // Ex: whats-hot
  const recordName = state + "wx";

  // A display name for your feed
  // Ex: What's Hot
  const displayName = state.toUpperCase() + "WX";

  // (Optional) A description of your feed
  // Ex: Top trending content from the whole network
  const description = "A feed containing all #" + state.toUpperCase() +
    "WX posts.";

  // (Optional) The path to an image to be used as your feed's avatar
  // Ex: ~/path/to/avatar.jpeg
  const avatar: string = "";

  // -------------------------------------
  // NO NEED TO TOUCH ANYTHING BELOW HERE
  // -------------------------------------

  if (
    !Deno.env.get("FEEDGEN_SERVICE_DID") && !Deno.env.get("FEEDGEN_HOSTNAME")
  ) {
    throw new Error("Please provide a hostname in the .env file");
  }
  const feedGenDid = Deno.env.get("FEEDGEN_SERVICE_DID") ??
    `did:web:${Deno.env.get("FEEDGEN_HOSTNAME")}`;

  // only update this if in a test environment
  const agent = new AtpAgent({ service: "https://bsky.social" });
  await agent.login({ identifier: handle, password });

  let avatarRef: BlobRef | undefined;
  if (avatar) {
    let encoding: string;
    if (avatar.endsWith("png")) {
      encoding = "image/png";
    } else if (avatar.endsWith("jpg") || avatar.endsWith("jpeg")) {
      encoding = "image/jpeg";
    } else {
      throw new Error("expected png or jpeg");
    }
    const img = await Deno.readFile(avatar);
    const blobRes = await agent.api.com.atproto.repo.uploadBlob(img, {
      encoding,
    });
    avatarRef = blobRes.data.blob;
  }

  await agent.api.com.atproto.repo.putRecord({
    repo: agent.session?.did ?? "",
    collection: ids.AppBskyFeedGenerator,
    rkey: recordName,
    record: {
      did: feedGenDid,
      displayName: displayName,
      description: description,
      avatar: avatarRef,
      createdAt: new Date().toISOString(),
    },
  });

  console.log(state + " added");
};

let i = 0;
const intervalId = setInterval(function () {
  if (i === states.length) {
    clearInterval(intervalId);
  }

  run(states[i]);

  i++;
}, 5000);

const ids = {
  ComAtprotoAdminDefs: "com.atproto.admin.defs",
  ComAtprotoAdminDisableAccountInvites:
    "com.atproto.admin.disableAccountInvites",
  ComAtprotoAdminDisableInviteCodes: "com.atproto.admin.disableInviteCodes",
  ComAtprotoAdminEnableAccountInvites: "com.atproto.admin.enableAccountInvites",
  ComAtprotoAdminGetInviteCodes: "com.atproto.admin.getInviteCodes",
  ComAtprotoAdminGetModerationAction: "com.atproto.admin.getModerationAction",
  ComAtprotoAdminGetModerationActions: "com.atproto.admin.getModerationActions",
  ComAtprotoAdminGetModerationReport: "com.atproto.admin.getModerationReport",
  ComAtprotoAdminGetModerationReports: "com.atproto.admin.getModerationReports",
  ComAtprotoAdminGetRecord: "com.atproto.admin.getRecord",
  ComAtprotoAdminGetRepo: "com.atproto.admin.getRepo",
  ComAtprotoAdminResolveModerationReports:
    "com.atproto.admin.resolveModerationReports",
  ComAtprotoAdminReverseModerationAction:
    "com.atproto.admin.reverseModerationAction",
  ComAtprotoAdminSearchRepos: "com.atproto.admin.searchRepos",
  ComAtprotoAdminSendEmail: "com.atproto.admin.sendEmail",
  ComAtprotoAdminTakeModerationAction: "com.atproto.admin.takeModerationAction",
  ComAtprotoAdminUpdateAccountEmail: "com.atproto.admin.updateAccountEmail",
  ComAtprotoAdminUpdateAccountHandle: "com.atproto.admin.updateAccountHandle",
  ComAtprotoIdentityResolveHandle: "com.atproto.identity.resolveHandle",
  ComAtprotoIdentityUpdateHandle: "com.atproto.identity.updateHandle",
  ComAtprotoLabelDefs: "com.atproto.label.defs",
  ComAtprotoLabelQueryLabels: "com.atproto.label.queryLabels",
  ComAtprotoLabelSubscribeLabels: "com.atproto.label.subscribeLabels",
  ComAtprotoModerationCreateReport: "com.atproto.moderation.createReport",
  ComAtprotoModerationDefs: "com.atproto.moderation.defs",
  ComAtprotoRepoApplyWrites: "com.atproto.repo.applyWrites",
  ComAtprotoRepoCreateRecord: "com.atproto.repo.createRecord",
  ComAtprotoRepoDeleteRecord: "com.atproto.repo.deleteRecord",
  ComAtprotoRepoDescribeRepo: "com.atproto.repo.describeRepo",
  ComAtprotoRepoGetRecord: "com.atproto.repo.getRecord",
  ComAtprotoRepoListRecords: "com.atproto.repo.listRecords",
  ComAtprotoRepoPutRecord: "com.atproto.repo.putRecord",
  ComAtprotoRepoStrongRef: "com.atproto.repo.strongRef",
  ComAtprotoRepoUploadBlob: "com.atproto.repo.uploadBlob",
  ComAtprotoServerCreateAccount: "com.atproto.server.createAccount",
  ComAtprotoServerCreateAppPassword: "com.atproto.server.createAppPassword",
  ComAtprotoServerCreateInviteCode: "com.atproto.server.createInviteCode",
  ComAtprotoServerCreateInviteCodes: "com.atproto.server.createInviteCodes",
  ComAtprotoServerCreateSession: "com.atproto.server.createSession",
  ComAtprotoServerDefs: "com.atproto.server.defs",
  ComAtprotoServerDeleteAccount: "com.atproto.server.deleteAccount",
  ComAtprotoServerDeleteSession: "com.atproto.server.deleteSession",
  ComAtprotoServerDescribeServer: "com.atproto.server.describeServer",
  ComAtprotoServerGetAccountInviteCodes:
    "com.atproto.server.getAccountInviteCodes",
  ComAtprotoServerGetSession: "com.atproto.server.getSession",
  ComAtprotoServerListAppPasswords: "com.atproto.server.listAppPasswords",
  ComAtprotoServerRefreshSession: "com.atproto.server.refreshSession",
  ComAtprotoServerRequestAccountDelete:
    "com.atproto.server.requestAccountDelete",
  ComAtprotoServerRequestPasswordReset:
    "com.atproto.server.requestPasswordReset",
  ComAtprotoServerResetPassword: "com.atproto.server.resetPassword",
  ComAtprotoServerRevokeAppPassword: "com.atproto.server.revokeAppPassword",
  ComAtprotoSyncGetBlob: "com.atproto.sync.getBlob",
  ComAtprotoSyncGetBlocks: "com.atproto.sync.getBlocks",
  ComAtprotoSyncGetCheckout: "com.atproto.sync.getCheckout",
  ComAtprotoSyncGetHead: "com.atproto.sync.getHead",
  ComAtprotoSyncGetLatestCommit: "com.atproto.sync.getLatestCommit",
  ComAtprotoSyncGetRecord: "com.atproto.sync.getRecord",
  ComAtprotoSyncGetRepo: "com.atproto.sync.getRepo",
  ComAtprotoSyncListBlobs: "com.atproto.sync.listBlobs",
  ComAtprotoSyncListRepos: "com.atproto.sync.listRepos",
  ComAtprotoSyncNotifyOfUpdate: "com.atproto.sync.notifyOfUpdate",
  ComAtprotoSyncRequestCrawl: "com.atproto.sync.requestCrawl",
  ComAtprotoSyncSubscribeRepos: "com.atproto.sync.subscribeRepos",
  AppBskyActorDefs: "app.bsky.actor.defs",
  AppBskyActorGetPreferences: "app.bsky.actor.getPreferences",
  AppBskyActorGetProfile: "app.bsky.actor.getProfile",
  AppBskyActorGetProfiles: "app.bsky.actor.getProfiles",
  AppBskyActorGetSuggestions: "app.bsky.actor.getSuggestions",
  AppBskyActorProfile: "app.bsky.actor.profile",
  AppBskyActorPutPreferences: "app.bsky.actor.putPreferences",
  AppBskyActorSearchActors: "app.bsky.actor.searchActors",
  AppBskyActorSearchActorsTypeahead: "app.bsky.actor.searchActorsTypeahead",
  AppBskyEmbedExternal: "app.bsky.embed.external",
  AppBskyEmbedImages: "app.bsky.embed.images",
  AppBskyEmbedRecord: "app.bsky.embed.record",
  AppBskyEmbedRecordWithMedia: "app.bsky.embed.recordWithMedia",
  AppBskyFeedDefs: "app.bsky.feed.defs",
  AppBskyFeedDescribeFeedGenerator: "app.bsky.feed.describeFeedGenerator",
  AppBskyFeedGenerator: "app.bsky.feed.generator",
  AppBskyFeedGetActorFeeds: "app.bsky.feed.getActorFeeds",
  AppBskyFeedGetActorLikes: "app.bsky.feed.getActorLikes",
  AppBskyFeedGetAuthorFeed: "app.bsky.feed.getAuthorFeed",
  AppBskyFeedGetFeed: "app.bsky.feed.getFeed",
  AppBskyFeedGetFeedGenerator: "app.bsky.feed.getFeedGenerator",
  AppBskyFeedGetFeedGenerators: "app.bsky.feed.getFeedGenerators",
  AppBskyFeedGetFeedSkeleton: "app.bsky.feed.getFeedSkeleton",
  AppBskyFeedGetLikes: "app.bsky.feed.getLikes",
  AppBskyFeedGetListFeed: "app.bsky.feed.getListFeed",
  AppBskyFeedGetPostThread: "app.bsky.feed.getPostThread",
  AppBskyFeedGetPosts: "app.bsky.feed.getPosts",
  AppBskyFeedGetRepostedBy: "app.bsky.feed.getRepostedBy",
  AppBskyFeedGetSuggestedFeeds: "app.bsky.feed.getSuggestedFeeds",
  AppBskyFeedGetTimeline: "app.bsky.feed.getTimeline",
  AppBskyFeedLike: "app.bsky.feed.like",
  AppBskyFeedPost: "app.bsky.feed.post",
  AppBskyFeedRepost: "app.bsky.feed.repost",
  AppBskyFeedSearchPosts: "app.bsky.feed.searchPosts",
  AppBskyFeedThreadgate: "app.bsky.feed.threadgate",
  AppBskyGraphBlock: "app.bsky.graph.block",
  AppBskyGraphDefs: "app.bsky.graph.defs",
  AppBskyGraphFollow: "app.bsky.graph.follow",
  AppBskyGraphGetBlocks: "app.bsky.graph.getBlocks",
  AppBskyGraphGetFollowers: "app.bsky.graph.getFollowers",
  AppBskyGraphGetFollows: "app.bsky.graph.getFollows",
  AppBskyGraphGetList: "app.bsky.graph.getList",
  AppBskyGraphGetListBlocks: "app.bsky.graph.getListBlocks",
  AppBskyGraphGetListMutes: "app.bsky.graph.getListMutes",
  AppBskyGraphGetLists: "app.bsky.graph.getLists",
  AppBskyGraphGetMutes: "app.bsky.graph.getMutes",
  AppBskyGraphGetSuggestedFollowsByActor:
    "app.bsky.graph.getSuggestedFollowsByActor",
  AppBskyGraphList: "app.bsky.graph.list",
  AppBskyGraphListblock: "app.bsky.graph.listblock",
  AppBskyGraphListitem: "app.bsky.graph.listitem",
  AppBskyGraphMuteActor: "app.bsky.graph.muteActor",
  AppBskyGraphMuteActorList: "app.bsky.graph.muteActorList",
  AppBskyGraphUnmuteActor: "app.bsky.graph.unmuteActor",
  AppBskyGraphUnmuteActorList: "app.bsky.graph.unmuteActorList",
  AppBskyNotificationGetUnreadCount: "app.bsky.notification.getUnreadCount",
  AppBskyNotificationListNotifications:
    "app.bsky.notification.listNotifications",
  AppBskyNotificationRegisterPush: "app.bsky.notification.registerPush",
  AppBskyNotificationUpdateSeen: "app.bsky.notification.updateSeen",
  AppBskyRichtextFacet: "app.bsky.richtext.facet",
  AppBskyUnspeccedDefs: "app.bsky.unspecced.defs",
  AppBskyUnspeccedGetPopular: "app.bsky.unspecced.getPopular",
  AppBskyUnspeccedGetPopularFeedGenerators:
    "app.bsky.unspecced.getPopularFeedGenerators",
  AppBskyUnspeccedGetTimelineSkeleton: "app.bsky.unspecced.getTimelineSkeleton",
  AppBskyUnspeccedSearchActorsSkeleton:
    "app.bsky.unspecced.searchActorsSkeleton",
  AppBskyUnspeccedSearchPostsSkeleton: "app.bsky.unspecced.searchPostsSkeleton",
};
